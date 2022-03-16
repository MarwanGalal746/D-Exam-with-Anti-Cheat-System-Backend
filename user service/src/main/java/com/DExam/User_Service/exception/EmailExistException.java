package com.DExam.User_Service.exception;

import com.DExam.User_Service.utility.CustomResponse;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Component;

@Component
public class EmailExistException extends RuntimeException{

    public ResponseEntity<?> getException(){
        return new ResponseEntity<>(
                new CustomResponse().setMessage(CustomResponse.EMAIL_USED).setStatus(HttpStatus.NOT_ACCEPTABLE)
                ,HttpStatus.NOT_ACCEPTABLE);
    }
}
