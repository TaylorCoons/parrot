package dev.joemedina.parrot.http;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.logging.Level;
import java.util.logging.Logger;

public class HTTPClient {
    private static final Logger logger = Logger.getLogger(HTTPClient.class.getName());
    private static final String API = "localhost:8080";
    private static final HttpClient HTTP_CLIENT = HttpClient.newBuilder().version(HttpClient.Version.HTTP_2).build();

    HTTPClient() {
        logger.log(Level.FINER, "Creating new HTTP2 client..");
    }

    private void createNewWorld(String worldName) {
        try {
            HttpRequest httpRequest = HttpRequest.newBuilder().POST(HttpRequest.BodyPublishers.ofString(worldName)).uri(URI.create(API+"/world")).build();
            HTTP_CLIENT.send(httpRequest, HttpResponse.BodyHandlers.ofString());
        }
        catch (InterruptedException e) {
            logger.log(Level.SEVERE, "Execution interrupted: {0}", e.getLocalizedMessage());
            Thread.currentThread().interrupt();
        }
        catch (IOException ioException) {
            logger.log(Level.SEVERE, "Error occurred when trying to create new world: {0}", ioException.getLocalizedMessage());
        }
    }


}
