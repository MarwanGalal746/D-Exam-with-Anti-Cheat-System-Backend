package dexam.desktopapplication.controllers;

import dexam.desktopapplication.Main;
import dexam.desktopapplication.api.ApiManager;
import dexam.desktopapplication.models.User;
import dexam.desktopapplication.utility.Camera;
import dexam.desktopapplication.utility.FileDownloader;
import javafx.fxml.FXML;
import javafx.scene.control.Button;
import javafx.scene.control.Label;
import javafx.scene.control.PasswordField;
import javafx.scene.control.TextField;

import java.io.IOException;
import java.net.URL;

public class LoginController {

    @FXML
    private Button loginButton;
    @FXML
    private Label errorText;
    @FXML
    private TextField emailField;
    @FXML
    private PasswordField passwordField;

    public void loginButtonClicked() throws IOException {

        if(emailField.getText().isEmpty() || passwordField.getText().isEmpty())
        {
            errorText.setText("Please enter your email and password");
            return;
        }
        boolean isValidUser = ApiManager.Login(new User(emailField.getText(), passwordField.getText()));
        //boolean isValidUser=true;
        if(!isValidUser) {
            errorText.setText("Wrong Credentials");
        } else {
            try{
                Camera.detect();
            }catch (Exception e){
                errorText.setText("Could not detect a camera");
                return;
            }
           try{
               URL url = new URL(Main.userImageURL);
               FileDownloader.downloadFile(url, "profilePicture.png");
           }catch (Exception e){
               errorText.setText("An Error has occurred while gathering your data, please try again");
               return;
              }
            Main.changeScene("browser.fxml");
        }
    }
}