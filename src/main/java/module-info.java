module dexam.desktopapplication {
    requires javafx.controls;
    requires javafx.fxml;
    requires unirest.java;
    requires webcam.capture;
    requires java.desktop;

    opens dexam.desktopapplication to javafx.fxml;
    exports dexam.desktopapplication;
    exports dexam.desktopapplication.controllers;
    opens dexam.desktopapplication.controllers to javafx.fxml;
    exports dexam.desktopapplication.api;
    opens dexam.desktopapplication.api to javafx.fxml;
    exports dexam.desktopapplication.anticheat;
    opens dexam.desktopapplication.anticheat to javafx.fxml;
    exports dexam.desktopapplication.models;
    opens dexam.desktopapplication.models to javafx.fxml;
    exports dexam.desktopapplication.utility;
    opens dexam.desktopapplication.utility to javafx.fxml;
}