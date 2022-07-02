package com.DExam.User_Service.model;

import com.DExam.User_Service.domain.Role;
import lombok.*;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class UserDTO {
    private String name;
    private String email;
    private String nationalID;
    private String password;
    private String img;
    private Role role;
}
