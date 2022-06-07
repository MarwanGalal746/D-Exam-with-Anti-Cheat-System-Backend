package com.DExam.User_Service.exception;

import com.DExam.User_Service.utility.CustomResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

public class EmailNotExistException extends RuntimeException{
    public ResponseEntity<?> getException(){
        return new ResponseEntity<>(
                new CustomResponse().setMessage(CustomResponse.INVALID_EMAIL).setStatus(HttpStatus.NOT_ACCEPTABLE)
                ,HttpStatus.NOT_ACCEPTABLE);
    }
}
