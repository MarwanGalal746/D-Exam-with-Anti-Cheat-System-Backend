package com.DExam.User_Service.resources.modules;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalTime;

@Entity
@Table(name = "users")
@Getter @Setter @NoArgsConstructor
public class User {
    private  @Id @GeneratedValue long id;
    private String name;
    private String email;
    private String nationalID;
    private String password;
    private String img;
    private Role role;
    private LocalTime createdAt;

    public User(String name, String email, String nationalID, String password, String img, Role role) {
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
        this.role = role;
        createdAt = LocalTime.now();
    }
}
