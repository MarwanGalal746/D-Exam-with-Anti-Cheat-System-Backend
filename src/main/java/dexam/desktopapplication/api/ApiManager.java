package dexam.desktopapplication.api;

import dexam.desktopapplication.Main;
import dexam.desktopapplication.models.User;
import kong.unirest.HttpResponse;
import kong.unirest.JsonNode;
import kong.unirest.Unirest;

import java.io.File;

public class ApiManager {
    public static boolean Login(User user) {
        HttpResponse<JsonNode> response = Unirest.post("http://20.125.84.58:8080/api/users/login")
                .header("Content-Type", "application/json")
                .body(user)
                .asJson();
        if (response.getStatus() == 200) {
            Main.userId = (int) response.getBody().getObject().getJSONObject("user").get("id");
            System.out.println(Main.userId);
            Main.userImageURL = (String) response.getBody().getObject().getJSONObject("user").get("img");
            System.out.println(Main.userImageURL);
            return true;
        } else {
            return false;
        }
    }

    public static boolean VerifyIdentity() {
        HttpResponse<JsonNode> response = Unirest.post("http://localhost:8000/verify/")
                .field("profilePicture", new File("./salah1.jpg"))
                .field("toComparePicture", new File("./profilePicture.png"))
                .asJson();
        return response.getStatus() == 200;
    }
}
