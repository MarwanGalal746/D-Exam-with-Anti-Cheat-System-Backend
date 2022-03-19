package com.DExam.User_Service.controller;

import com.DExam.User_Service.model.MailForm;
import com.DExam.User_Service.service.EmailService;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@AllArgsConstructor
@RestController
@RequestMapping("/api/mail")
public class EmailController {

    private final EmailService emailSender;

    @PostMapping("/send")
    public boolean send(@RequestBody MailForm mailForm) {

        emailSender.send(mailForm.getTo(),mailForm.getSubject(),mailForm.getBody());
        return true;
    }

}
