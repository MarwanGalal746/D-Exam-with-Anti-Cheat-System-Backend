package com.DExam.User_Service.resources.modules;

import com.DExam.User_Service.resources.modules.Role;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "users")
public class User {
    private  @Id @GeneratedValue int id;
    private String name;
    private String email;
    private String nationalID;
    private String password;
    private String img;
    private Role role;

    public User() {}

    public User(String name, String email, String nationalID, String password, String img, Role role) {
        this.name = name;
        this.email = email;
        this.nationalID = nationalID;
        this.password = password;
        this.img = img;
        this.role = role;
    }

    public int getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getNationalID() {
        return nationalID;
    }

    public void setNationalID(String nationalID) {
        this.nationalID = nationalID;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getImg() {
        return img;
    }

    public void setImg(String img) {
        this.img = img;
    }

    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }
}
