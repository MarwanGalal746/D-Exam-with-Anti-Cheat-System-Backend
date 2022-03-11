package com.DExam.User_Service.resources.modules;

import lombok.;

import javax.persistence.;
import java.text.SimpleDateFormat;
import java.time.LocalTime;
import java.util.Date;

@Entity
@Table(name = "users")
@Getter @Setter @NoArgsConstructor
public class User {
    private  @Id @GeneratedValue @Setter(AccessLevel.PROTECTED) long id;
    private String name;
    private String email;
    private String nationalID;
    private String password;
    private String img;
    private Role role;
    private @Setter(AccessLevel.PROTECTED) String createdAt;

    public User(String name, String email, String nationalID, String password, String img, Role role) {
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
        this.role = role;
        SimpleDateFormat formatter= new SimpleDateFormat("yyyy-MM-dd 'at' HH:mm:ss z");
        Date date = new Date(System.currentTimeMillis());
        System.out.println(formatter.format(date));
        createdAt = formatter.format(date);
    }
}
