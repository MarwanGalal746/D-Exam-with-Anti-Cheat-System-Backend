package com.DExam.User_Service.utility;

import lombok.Getter;
import org.springframework.http.HttpStatus;

@Getter
public class CustomResponse implements ResponseMessages{

    private String message;
    private String status;
    private String code;

    public CustomResponse setMessage(String message) {
        this.message = message;
        return this;
    }

    public CustomResponse setStatus(HttpStatus status) {
        this.status = status.getReasonPhrase();
        this.code = String.valueOf(status.value());
        return this;
    }
}
