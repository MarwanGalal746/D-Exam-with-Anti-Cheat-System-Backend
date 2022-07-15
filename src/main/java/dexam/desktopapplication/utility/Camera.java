package dexam.desktopapplication.utility;

import com.github.sarxos.webcam.Webcam;
import com.github.sarxos.webcam.WebcamException;

import java.io.IOException;

public class Camera {
    public static void detect() throws IOException {
        try{
            Webcam webcam = Webcam.getDefault();
        } catch (Exception e) {
            throw new WebcamException(e);
        }
    }
}