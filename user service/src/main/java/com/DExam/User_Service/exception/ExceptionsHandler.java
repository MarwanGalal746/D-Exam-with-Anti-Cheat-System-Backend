package com.DExam.User_Service.exception;

import com.DExam.User_Service.utility.CustomResponse;
import com.DExam.User_Service.utility.ResponseMessages;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import java.util.HashMap;

@ControllerAdvice
public class ExceptionsHandler {

    private final ObjectMapper objectMapper = new ObjectMapper();
    private final CustomResponse customResponse = new CustomResponse();
    private HashMap response;

    @ExceptionHandler(UserNotFoundException.class)
    public ResponseEntity<?> userNotFound(){
        CustomResponse customResponse = new CustomResponse();
        customResponse.setMessage(ResponseMessages.USER_NOT_FOUND).setStatus(HttpStatus.NOT_FOUND);
        response = objectMapper.convertValue(customResponse,HashMap.class);
        return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
    }

    @ExceptionHandler(EmailExistException.class)
    public ResponseEntity<?> emailExist(){
        customResponse.setMessage(ResponseMessages.EMAIL_USED).setStatus(HttpStatus.NOT_ACCEPTABLE);
        response = objectMapper.convertValue(customResponse,HashMap.class);
        return new ResponseEntity<>(response, HttpStatus.NOT_ACCEPTABLE);
    }

    @ExceptionHandler(NationalIDException.class)
    public ResponseEntity<?> nationalIdExist(){
        customResponse.setMessage(ResponseMessages.NATIONAL_ID_USED).setStatus(HttpStatus.NOT_ACCEPTABLE);
        response = objectMapper.convertValue(customResponse,HashMap.class);
        return new ResponseEntity<>(response, HttpStatus.NOT_ACCEPTABLE);
    }

    @ExceptionHandler(InvalidEmailPasswordException.class)
    public ResponseEntity<?> invalidEmailPassword(){
        customResponse.setMessage(ResponseMessages.INVALID_EMAIL_PASSWORD).setStatus(HttpStatus.NOT_FOUND);
        response = objectMapper.convertValue(customResponse,HashMap.class);
        return new ResponseEntity<>(response, HttpStatus.NOT_FOUND);
    }
}
