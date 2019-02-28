use crate::lexer::Token;

#[derive(Debug, PartialEq, Clone)]
pub enum SyntaxItem {
    // Branch nodes
    Operator(String),

    // Leaf nodes
    Integer(i64),
    DefaultASTNode,
}

impl SyntaxItem {
    pub fn is_plus_minus(self) -> bool {
        self == SyntaxItem::Operator("+".to_string()) ||
            self == SyntaxItem::Operator("-".to_string())
    }
}

#[derive(Debug, Clone)]
pub struct ASTNode {
    pub children: Vec<ASTNode>,
    pub item: SyntaxItem,
}

impl ASTNode {
    pub fn new() -> ASTNode {
        ASTNode {
            children: Vec::new(),
            item: SyntaxItem::DefaultASTNode,
        }
    }

    pub fn travel(self) -> ASTNode {
        match self.item {
            SyntaxItem::Integer(_) => self,
            SyntaxItem::Operator(op) => {
                if vec!["+", "-", "*", "/"].contains(&op.as_str()) {
                    if self.children.len() != 2 {
                        panic!("operator {} requires 2 operands", op)
                    }

                    let mut c = self.children.clone();
                    let right = c.pop().unwrap().travel();
                    let left = c.pop().unwrap().travel();

                    let mut node = ASTNode::new();
                    match (left.item, right.item) {
                        (SyntaxItem::Integer(l), SyntaxItem::Integer(r)) => {
                            match op.as_str() {
                                "+" => node.item = SyntaxItem::Integer(l + r),
                                "-" => node.item = SyntaxItem::Integer(l - r),
                                "*" => node.item = SyntaxItem::Integer(l * r),
                                "/" => node.item = SyntaxItem::Integer(l / r),
                                _ => {},
                            }
                        },
                        _ => {},
                    }
                    node
                } else {
                    panic!("not implemented operator: {}", op)
                }
            },
            _ => panic!("unexpected token"),
        }        
    }

    fn get_child_at(self, idx: usize) -> ASTNode {
        self.children.clone()[idx].clone()
    }

    fn rotate_left(self) -> ASTNode {
        let mut children = self.children.clone();
        let mut right = children.pop().unwrap();
        let right_right = right.children.pop().unwrap();
        let right_left = right.children.pop().unwrap();
        let left = children.pop().unwrap();

        ASTNode {
            item: right.item,
            children: vec![
                ASTNode {
                    item: self.item,
                    children: vec![left, right_right],
                },
                right_left,
            ]
        }
    }

    pub fn print(self, depth: usize) {
        println!("{:?}", self.item);
        for child in self.children {
            print!("{}", "  ".repeat(depth + 1));
            child.print(depth + 1);
        }
    }
}

pub struct Parser {
    tokens: Vec<Token>,
}

// expr   -> term (('+'|'-') expr)*
// term   -> factor (('*' | '/') factor)*
// factor -> INTEGER
impl Parser {
    pub fn new(t: Vec<Token>) -> Parser {
        Parser { tokens: t }
    }

    pub fn parse(&mut self) -> Result<ASTNode, String> {
        self.parse_expr(0).and_then(|(n, i)| 
            if i == self.tokens.len() {
                Ok(n)
            } else {
                Err(format!("Expects end of input, found {:?}", self.tokens[i]))
            }
        )
    }

    // expr -> term (('+'|'-') expr)*
    fn parse_expr(&mut self, pos: usize) -> Result<(ASTNode, usize), String> {
        match self.parse_term(pos) {
            Ok((term, next_pos)) => {
                match self.tokens.get(next_pos) {
                    Some(Token::Operator(op)) 
                    if op == &"+" || op == &"-" => {
                        let mut node = ASTNode::new();
                        node.item = SyntaxItem::Operator(op.to_string());
                        node.children.push(term);

                        match self.parse_expr(next_pos + 1) {
                            Ok((expr, expr_next_pos)) => {
                                node.children.push(expr.clone());

                                // Need to rotate tree for operators with same priority.
                                if expr.item.is_plus_minus() {
                                    node = node.rotate_left();
                                }
                                    
                                Ok((node, expr_next_pos))
                            },
                            Err(e) => Err(e),
                        }
                    },
                    _ => Ok((term, next_pos)),
                }
            },
            Err(e) => Err(e),
        }
        
    }

    // term -> factor (('*' | '/') factor)*
    fn parse_term(&mut self, pos: usize) -> Result<(ASTNode, usize), String> {
        match self.parse_factor(pos) {
            Ok((factor, next_pos)) => {
                match self.tokens.get(next_pos) {
                    Some(Token::Operator(op))
                    if op == &"*" || op == &"/" => {
                        let mut node = ASTNode::new();
                        node.item = SyntaxItem::Operator(op.to_string());
                        node.children.push(factor);

                        match self.parse_term(next_pos + 1) {
                            Ok((expr, expr_next_pos)) => {
                                node.children.push(expr);
                                Ok((node, expr_next_pos))
                            },
                            Err(e) => Err(e),
                        }
                    },
                    _ => Ok((factor, next_pos)),
                }
            },
            Err(e) => Err(e),
        }
    }

    // factor -> INTEGER
    fn parse_factor(&mut self, pos: usize) -> Result<(ASTNode, usize), String> {
        match self.tokens.get(pos) {
            Some(token) => match token {
                Token::Integer(n) => {
                    let mut node = ASTNode::new();
                    node.item = SyntaxItem::Integer(n.parse::<i64>().unwrap());
                    Ok((node, pos + 1))
                },
                _ => Err(format!("Unexpected token {:?}", { token })),
            },
            None => return Err("Unexpected end of input".to_string()),
        }
    }
}

mod tests {
    use super::*;

    // 10 + 3
    #[test]
    fn test_parser_1() {
        let tokens = vec![
            Token::Integer("10".to_string()),
            Token::Operator("+".to_string()),
            Token::Integer("3".to_string()),
        ];

        let mut parser = Parser::new(tokens);
        match parser.parse() {
            Ok(node) => {
                assert_eq!(node.item, SyntaxItem::Operator("+".to_string()));
                assert_eq!(node.children[0].item, SyntaxItem::Integer(10));
                assert_eq!(node.children[1].item, SyntaxItem::Integer(3));
            },
            Err(e) => panic!(e),
        }
    }

    // 10 * 4 + 2
    #[test]
    fn test_parser_2() {
        let tokens = vec![
            Token::Integer("10".to_string()),
            Token::Operator("*".to_string()),
            Token::Integer("4".to_string()),
            Token::Operator("+".to_string()),
            Token::Integer("2".to_string()),
        ];

        let mut parser = Parser::new(tokens);
        match parser.parse() {
            Ok(node) => {
                assert_eq!(node.item, SyntaxItem::Operator("+".to_string()));
                assert_eq!(node.children[0].item, SyntaxItem::Operator("*".to_string()));
                assert_eq!(node.children[0].children[0].item, SyntaxItem::Integer(10));
                assert_eq!(node.children[0].children[1].item, SyntaxItem::Integer(4));
                assert_eq!(node.children[1].item, SyntaxItem::Integer(2));
            },
            Err(e) => panic!(e),
        }
    }

    // 2 + 10 * 4
    #[test]
    fn test_parser_3() {
        let tokens = vec![
            Token::Integer("2".to_string()),
            Token::Operator("+".to_string()),
            Token::Integer("10".to_string()),
            Token::Operator("*".to_string()),
            Token::Integer("4".to_string()),
        ];

        let mut parser = Parser::new(tokens);
        match parser.parse() {
            Ok(node) => {
                assert_eq!(node.item, SyntaxItem::Operator("+".to_string()));
                assert_eq!(node.children[0].item, SyntaxItem::Integer(2));
                assert_eq!(node.children[1].item, SyntaxItem::Operator("*".to_string()));
                assert_eq!(node.children[1].children[0].item, SyntaxItem::Integer(10));
                assert_eq!(node.children[1].children[1].item, SyntaxItem::Integer(4));
            },
            Err(e) => panic!(e),
        }
    }
}
