import java.io.Serializable;

public class Hello implements Serializable {
  public static final int MAGIC = 42;

  private String message;

  public Hello(String message) {
    this.message = message;
  }

  public static void main(String[] args) {
    System.out.println("Hello world");
  }
}
