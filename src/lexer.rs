use std::str::Chars;
use std::iter::Peekable;

#[derive(Debug, PartialEq)]
pub enum Token {
    // Literals
    Integer(String),

    // Operators
    Operator(String),

    // Non-codes
    Comment,
    EOF,
    Illegal,
}

pub struct Lexer<'a> {
    input: Peekable<Chars<'a>>,
}

impl<'a> Lexer<'a> {
    pub fn new(input: &str) -> Lexer {
        Lexer { input: input.chars().peekable() }
    }

    pub fn next_token(&mut self) -> Token {
        self.skip_whitespace();

        match self.input.next() {
            Some('+') => Token::Operator("+".to_string()),
            Some('-') => Token::Operator("-".to_string()),
            Some('*') => Token::Operator("*".to_string()),
            Some('/') => Token::Operator("/".to_string()),
            Some(ch @ _) => {
                if ch.is_numeric() {
                    let mut number = String::new();
                    number.push(ch);
                    while let Some(&c) = self.input.peek() {
                        if !c.is_numeric() {
                            break;
                        }
                        number.push(self.input.next().unwrap());
                    }

                    Token::Integer(number)
                } else {
                    Token::Illegal
                }
            }
            None => Token::EOF,
        }
    }

    fn skip_whitespace(&mut self) {
        while let Some(&c) = self.input.peek() {
            if !c.is_whitespace() {
                break;
            }
            self.input.next();
        }
    }
}

#[test]
fn test_lexer() {
    let input = "10 + 3";
    let expected = vec![
        Token::Integer("10".to_string()),
        Token::Operator("+".to_string()),
        Token::Integer("3".to_string()),
    ];

    let mut lexer = Lexer::new(input);
    for t in expected {
        let token = lexer.next_token();
        assert_eq!(token, t);
    }
}
