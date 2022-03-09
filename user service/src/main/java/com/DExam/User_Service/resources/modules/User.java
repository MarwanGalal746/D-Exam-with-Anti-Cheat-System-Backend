package com.DExam.User_Service.resources.modules;

import com.DExam.User_Service.resources.modules.Role;
import lombok.Data;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalTime;
import java.util.Date;
import java.util.TimeZone;

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
    @OneToOne
    @JoinColumn(name = "role_id")
    private Role role;
    @Transient
    private LocalTime createdAt;

    public User(long id, String name, String email, String nationalID, String password, String img, Role role) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
        this.role = role;
        createdAt = LocalTime.now();
    }
}
