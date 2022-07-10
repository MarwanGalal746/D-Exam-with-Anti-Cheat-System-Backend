package dexam.desktopapplication.anticheat;

import dexam.desktopapplication.api.ApiManager;
import dexam.desktopapplication.utility.Camera;

import java.io.IOException;
import java.util.ArrayList;
import java.util.concurrent.TimeUnit;

public class AntiCheat {
    public static void initiate(String browser) throws IOException, InterruptedException {

       runProcessManager(browser);

        Thread newThread2 = new Thread(()->{
            while (true){
                try {
                    TimeUnit.SECONDS.sleep(60);
                    Camera.detect();
                    ApiManager.VerifyIdentity();
                } catch (IOException | InterruptedException e) {
                    throw new RuntimeException(e);
                }
            }
        });
        newThread2.start();

        return;
    }

    private static void runProcessManager(String browser) throws IOException, InterruptedException {
        ProcessManager processManager = new ProcessManager();

        ArrayList<String> p = new ArrayList<>();
        p.add("java.exe");
        p.add("javaw.exe");
        p.add("Discord.exe");
        p.add("idea64.exe");
        p.add("Music.UI.exe");
        p.add("Docker Desktop.exe");
        p.add("docker.exe");
        p.add(browser);

        processManager.addToIgnoreList(p);
        processManager.killProcesses();
        Thread newThread = new Thread(() -> {
            try {
                processManager.monitorActivity();
            } catch (IOException e) {
                throw new RuntimeException(e);
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
        });
        newThread.start();
    }
}
