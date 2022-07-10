package dexam.desktopapplication.controllers;

import dexam.desktopapplication.Main;
import dexam.desktopapplication.anticheat.AntiCheat;
import dexam.desktopapplication.api.ApiManager;
import javafx.fxml.Initializable;

import java.io.IOException;
import java.net.URL;
import java.util.ResourceBundle;

public class StartAppController implements Initializable {
    @Override
    public void initialize(URL url, ResourceBundle resourceBundle) {
        try {
            AntiCheat.initiate(Main.usedBrowser);
            System.out.println(ApiManager.VerifyIdentity());
            Main.changeScene("stopApp.fxml");
        } catch (IOException e) {
            throw new RuntimeException(e);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
    }
}
