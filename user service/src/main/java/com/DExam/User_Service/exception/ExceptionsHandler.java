package com.DExam.User_Service.exception;

import com.DExam.User_Service.utility.CustomResponse;
import com.DExam.User_Service.utility.ResponseMessages;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;

@ControllerAdvice
public class ExceptionsHandler {

    private final CustomResponse response = new CustomResponse();

    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<?> userNotFound(){
        response.setMessage(ResponseMessages.USER_NOT_FOUND).setStatus(HttpStatus.NOT_FOUND);
        return new ResponseEntity<>(response,HttpStatus.NOT_FOUND);
    }

    @ExceptionHandler(EmailExistException.class)
    public ResponseEntity<?> emailExist(){
        response.setMessage(ResponseMessages.EMAIL_USED).setStatus(HttpStatus.NOT_ACCEPTABLE);
        return new ResponseEntity<>(response, HttpStatus.NOT_ACCEPTABLE);
    }

    @ExceptionHandler(NationalIDException.class)
    public ResponseEntity<?> nationalIdExist(){
        response.setMessage(ResponseMessages.NATIONAL_ID_USED).setStatus(HttpStatus.NOT_ACCEPTABLE);
        return new ResponseEntity<>(response, HttpStatus.NOT_ACCEPTABLE);
    }

    @ExceptionHandler(InvalidEmailPasswordException.class)
    public ResponseEntity<?> invalidEmailPassword(){
        response.setMessage(ResponseMessages.INVALID_EMAIL_PASSWORD).setStatus(HttpStatus.NOT_FOUND);
        return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
    }
}
