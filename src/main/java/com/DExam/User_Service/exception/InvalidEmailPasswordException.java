package com.DExam.User_Service.exception;

import com.DExam.User_Service.utility.CustomResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;

@Component
public class InvalidEmailPasswordException extends RuntimeException{

    public ResponseEntity<?> getException(){
        return new ResponseEntity<>(
                new CustomResponse().setMessage(CustomResponse.INVALID_EMAIL_PASSWORD).setStatus(HttpStatus.NOT_FOUND)
                ,HttpStatus.NOT_FOUND);
    }
}
