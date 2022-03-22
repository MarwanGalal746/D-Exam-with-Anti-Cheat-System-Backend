package com.DExam.User_Service.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data @NoArgsConstructor @AllArgsConstructor
public class UpdateUserRequest {
    private User oldUser;
    private User newUser;
}
