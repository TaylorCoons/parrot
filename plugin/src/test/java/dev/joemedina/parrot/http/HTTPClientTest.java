package dev.joemedina.parrot.http;

import org.apache.http.client.methods.HttpDelete;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.mockserver.integration.ClientAndServer;
import org.mockserver.socket.PortFactory;
import org.mockserver.verify.VerificationTimes;

import java.net.URI;

import static org.mockserver.model.HttpRequest.request;
import static org.mockserver.model.StringBody.exact;

public class HTTPClientTest {

    private ClientAndServer mockServer;
    private URI testUri;

    @Before
    public void startMockServer() {
        int port = PortFactory.findFreePort();
        this.testUri = URI.create("http://localhost:" + port);
        mockServer = ClientAndServer.startClientAndServer(port);
    }

    @Test
    public void createNewWorldTest() {
        HTTPClient httpClient = new HTTPClient(testUri);
        httpClient.createNewWorld("TestWorld");
        mockServer.verify(
                request()
                .withMethod(HttpPost.METHOD_NAME)
                .withPath("/world")
                .withBody(exact("TestWorld")),
                VerificationTimes.exactly(1)
        );
    }

    @Test
    public void getWorldListTest() {
        HTTPClient httpClient = new HTTPClient(testUri);
        httpClient.getWorldList();
        mockServer.verify(
                request()
                .withMethod(HttpGet.METHOD_NAME)
                .withPath("/world"),
        VerificationTimes.exactly(1)
        );
    }

    @Test
    public void deleteWorldTest() {
        HTTPClient httpClient = new HTTPClient(testUri);
        httpClient.deleteWorld("TestWorld");
        mockServer.verify(
                request()
                .withMethod(HttpDelete.METHOD_NAME)
                .withPath("/world/TestWorld"),
                VerificationTimes.exactly(1)
        );
    }

    @After
    public void stopMockServer() {
        mockServer.stop();
    }
}
