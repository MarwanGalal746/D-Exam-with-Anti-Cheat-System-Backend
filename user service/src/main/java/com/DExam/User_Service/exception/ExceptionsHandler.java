package com.DExam.User_Service.exception;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

@ControllerAdvice
public class ExceptionsHandler {

    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<?> userNotFound(UserNotFoundException exception){
        return exception.getException();
    }

    @ExceptionHandler(EmailExistException.class)
    public ResponseEntity<?> emailExist(EmailExistException exception){
        return exception.getException();
    }

    @ExceptionHandler(NationalIDException.class)
    public ResponseEntity<?> nationalIdExist(NationalIDException exception){
        return exception.getException();
    }

    @ExceptionHandler(InvalidEmailPasswordException.class)
    public ResponseEntity<?> invalidEmailPassword(InvalidEmailPasswordException exception){
        return exception.getException();
    }

    @ExceptionHandler(EmailNotExistException.class)
    public ResponseEntity<?> emailNotExist(EmailNotExistException exception){
        return exception.getException();
    }
}
