pub mod lexer;
pub mod parser;

use std::io;
use std::io::prelude::*;
use crate::lexer::{Lexer, Token};
use crate::parser::{Parser, ASTNode};


fn lex(input: &str) -> Vec<Token> {
    let mut lexer = Lexer::new(input);
    let mut tokens = Vec::new();
    loop {
        let token = lexer.next_token();
        if token == Token::EOF {
            break;
        }
        tokens.push(token);
    }

    tokens
}

fn parse<'a>(tokens: Vec<Token>) -> ASTNode {
    let mut parser = Parser::new(tokens);
    match parser.parse() {
        Ok(root) => root,
        Err(message) => {
            panic!(message)
        },
    }
}

fn interpret() {
    print!("\n>>> ");
    io::stdout().flush();

    let mut input = String::new();
    let result = io::stdin().read_line(&mut input);
    if result.is_err() {
        println!("input error: {:?}", result.err())
    }
    
    let tokens = lex(input.as_str());
    let root = parse(tokens);
    root.clone().print(0);
    println!("{:?}", root.travel().item);
}

fn main() {
    loop {
        interpret();
    }
}
