package com.DExam.User_Service.controller;

import com.DExam.User_Service.config.JwtManager;
import com.DExam.User_Service.service.UserService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.http.MediaType;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(UserController.class)
class UserControllerTest {

    @MockBean
    private UserService userService;
    @MockBean
    private JwtManager jwtManager;
    @MockBean
    private AuthenticationManager authenticationManager;
    @MockBean
    private EmailController emailController;

    @Autowired
    private MockMvc mockMvc;



    @Test
    void get() throws Exception {

    }

    @Test
    void verify() {
    }

    @Test
    void register() throws Exception {
//        long id =1;
//        User user = new User(1,"Ahmed Mohamed", "eyaaaada@gmail.com", "30001dsdeqaq3aaaaaaaaaa69666",
//                Role.Student, "00",  "123123123123124jw33o22qewq",
//                new Date(2022,03,20 ,16,10,23));
//        when(userService.add(user)).thenReturn(id);

        mockMvc.perform(MockMvcRequestBuilders.post("/api/users/register")
                        .contentType(MediaType.APPLICATION_JSON)
                        .content("{\n" +
                                "  \"name\": \"Ahmed Mohamed\",\n" +
                                "  \"email\": \"eyad@gmail.com\",\n" +
                                "  \"nationalID\": \"30001dsdeqq369666\",\n" +
                                "  \"img\": \"123123123123124jw33o22qewq\",\n" +
                                "  \"role\":\"Student\"\n" +
                                "}"))
                .andExpect(status().isCreated());
    }

    @Test
    void update() {
    }

    @Test
    void login() {
    }
}