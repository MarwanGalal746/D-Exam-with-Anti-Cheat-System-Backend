package com.DExam.User_Service.resources.modules;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import java.text.SimpleDateFormat;
import java.util.Date;

@Entity
@Table(name = "users")
@Getter @Setter @NoArgsConstructor
public class User {
    private  @Id @GeneratedValue @Setter(AccessLevel.PROTECTED) long id;
    private String name;
    @Column(unique=true)
    private String email;
    @Column(unique=true)
    private String nationalID;
    private String password;
    private String img;
    private @Setter(AccessLevel.PROTECTED) String createdAt;

    public User(String name, String email, String nationalID, String password, String img, Role role) {
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
        SimpleDateFormat formatter= new SimpleDateFormat("yyyy-MM-dd 'at' HH:mm:ss z");
        Date date = new Date(System.currentTimeMillis());
        System.out.println(formatter.format(date));
        createdAt = formatter.format(date);
    }
}
