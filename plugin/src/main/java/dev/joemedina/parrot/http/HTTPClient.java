package dev.joemedina.parrot.http;

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.Collections;
import java.util.logging.Level;
import java.util.logging.Logger;

public class HTTPClient {
    private static final Logger logger = Logger.getLogger(HTTPClient.class.getName());
    private static final HttpClient HTTP_CLIENT = HttpClient.newBuilder().version(HttpClient.Version.HTTP_2).build();
    private static final String WORLD_LOCATION = "/world";
    private final URI apiUri;

    HTTPClient(URI uri) {
        this.apiUri = uri;
        logger.log(Level.FINER, "Creating new HTTP client with URI: {0}", Collections.singletonList(apiUri));
    }

    private String performRestCall(HttpRequest httpRequest) {
        String response = null;
        try {
            response = HTTP_CLIENT.send(httpRequest, HttpResponse.BodyHandlers.ofString()).body();
        }
        catch (InterruptedException e) {
            logger.log(Level.SEVERE, "Execution interrupted: {0}", e.getLocalizedMessage());
            Thread.currentThread().interrupt();
        }
        catch (IOException ioException) {
            logger.log(Level.SEVERE, "Error occurred when trying to create new world: {0}", new String[] {ioException.toString()});
        }
        return response;
    }

    void createNewWorld(String worldName) {
        HttpRequest httpRequest = HttpRequest.newBuilder().POST(HttpRequest.BodyPublishers.ofString(worldName)).uri(URI.create(apiUri +WORLD_LOCATION)).build();
        performRestCall(httpRequest);
    }

    void getWorldList() {
        HttpRequest httpRequest = HttpRequest.newBuilder().GET().uri(URI.create(apiUri + WORLD_LOCATION)).build();
        performRestCall(httpRequest);
    }

    void deleteWorld(String worldName) {
        HttpRequest httpRequest = HttpRequest.newBuilder().DELETE().uri(URI.create(apiUri + WORLD_LOCATION + "/" + worldName)).build();
        performRestCall(httpRequest);
    }
}
