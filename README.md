# TGSender
Service for answering questions by mail, using a telegram bot

## Install

1. Clone this repository
2. Edit config in config/confim.yml for your settings (bot token)
3. Input command 'make compose-up' or 'make compose-up-silent'(silent mode) - up server, telegram bot, postgres DB

## Use
1. Input localhost:8080/question in your browser for Ask a Question.
2. Input localhost:8080/question/{id} for show your question
3. Input localhost:8080/question/all for show all questions *(not implements)
4. Input message '/all' for show all Open question *(not implements)