package dexam.desktopapplication.controllers;

import dexam.desktopapplication.Main;
import javafx.fxml.FXML;
import javafx.scene.control.Button;

public class StopAppController {
    @FXML
    public Button stopButton;

    public void stopAntiCheat() {
        Main.stopApp();
    }
}
