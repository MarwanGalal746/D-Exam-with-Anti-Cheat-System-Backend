package dexam.desktopapplication;

import dexam.desktopapplication.messaging.RabbitMq;
import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.stage.Stage;

import java.io.IOException;
import java.util.Objects;

public class Main extends Application {

    private static Stage currentStage;
    public static int userId;
    public static String usedBrowser;
    public static String userImageURL;
    public static void stopApp() {
        currentStage.close();
    }

    @Override
    public void start(Stage stage) throws IOException {
        currentStage = stage;
        Parent root = FXMLLoader.load(Objects.requireNonNull(getClass().getResource("stopApp.fxml")));
        Scene scene = new Scene(root);
        stage.setTitle("Anti-Cheat");
        stage.setScene(scene);
        stage.setResizable(false);
        stage.show();
    }

    @Override
    public void stop() throws IOException {
        System.out.println("Stage is closing");
        RabbitMq.send("close-" + userId);
        String[] cmd = {"cmd.exe", "/c", "start explorer.exe"};
        Runtime.getRuntime().exec(cmd);
    }

    public static void changeScene(String fxmlFile) throws IOException {
        Parent pane = FXMLLoader.load(Objects.requireNonNull(Main.class.getResource(fxmlFile)));
        Scene scene = new Scene(pane);
        currentStage.setScene(scene);
        currentStage.show();
    }

    public static void main(String[] args) {
        launch();
    }
}