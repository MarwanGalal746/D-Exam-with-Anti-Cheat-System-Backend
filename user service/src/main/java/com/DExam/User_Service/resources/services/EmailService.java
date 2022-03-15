package com.DExam.User_Service.resources.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

@Service
public class EmailService{

    @Autowired
    private JavaMailSender mailSender;

    public EmailService() {
    }

    public void send(String to, String subject, String emailBody) {
        SimpleMailMessage msg = new SimpleMailMessage();
        msg.setTo(to);
        msg.setText(emailBody);
        msg.setSubject(subject);
        mailSender.send(msg);
    }
}
