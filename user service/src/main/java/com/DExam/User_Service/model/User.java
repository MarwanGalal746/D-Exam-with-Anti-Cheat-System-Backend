package com.DExam.User_Service.model;

import lombok.AccessLevel;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import javax.persistence.*;
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
    @Temporal(TemporalType.TIMESTAMP) @Column(nullable = false)
    private @Setter(AccessLevel.PROTECTED) Date createdAt;

    @PrePersist
    private void onCreation(){
        createdAt = new Date();
    }

    public User(String name, String email, String nationalID, String password, String img) {
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
    }
}
