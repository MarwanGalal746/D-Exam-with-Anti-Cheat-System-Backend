package com.DExam.User_Service.resources.modules;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table
public enum Role {

    Student,
    Teacher,
    ;
    private  @Id @GeneratedValue long id;
}
