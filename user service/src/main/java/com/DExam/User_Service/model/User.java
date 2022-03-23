package com.DExam.User_Service.model;

import lombok.*;
import org.hibernate.annotations.DynamicUpdate;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "users")
@DynamicUpdate
@Getter @Setter @NoArgsConstructor @AllArgsConstructor
public class User {
    private  @Id @GeneratedValue @Setter(AccessLevel.PROTECTED) long id;
    private String name;
    @Column(unique=true)
    private String email;
    @Column(unique=true)
    private String nationalID;
    @Enumerated(EnumType.STRING)
    private Role role;
    private String password;
    private String img;
    @Temporal(TemporalType.TIMESTAMP) @Column(nullable = false)
    private @Setter(AccessLevel.PROTECTED) Date createdAt;

    @PrePersist
    private void onCreation(){
        createdAt = new Date();
    }
}
