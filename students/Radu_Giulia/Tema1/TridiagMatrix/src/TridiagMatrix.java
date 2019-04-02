import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Scanner;

public class TridiagMatrix {

	public static double a[],b[],c[],r[],s[];
	public static double x[];
	public static double d[];
	public static int n;

	public static Scanner sc;

	static void citire() throws IOException
	{
		sc = new Scanner(new FileReader(new File("date.in")));
		n=sc.nextInt();
		x=new double[n+1];
		a=new double[n+1]; b=new double[n+1]; c=new double[n+1]; d=new double[n+1]; r=new double[n+1]; s=new double[n+1];

		for(int i=1;i<=n;i++) {
			a[i]=sc.nextDouble();
		}
		for(int i=1;i<=n-1;i++) {
			b[i]=sc.nextDouble();
		}
		for(int i=1;i<=n-1;i++) {
			c[i]=sc.nextDouble();
		}
		for(int i=1;i<=n;i++) {
			d[i]=sc.nextDouble();
		}


	}




	public static void main(String[] args) {

		try {
			citire();
		} catch (IOException e) {
			e.printStackTrace();
		}
		
		r[1]=-c[1]/a[1];
		s[1]=d[1]/a[1];
	
		for(int i=2;i<n;i++) {
			r[i]=-c[i]/(b[i]*r[i-1]+a[i]);
			s[i]=(d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
		}

		x[n]=(d[n] - b[n-1] * s[n-1]) / (b[n-1] * r[n-1] + a[n]);
		
		for(int i=n-1;i>=1;i--)
			x[i]= r[i] * x[i+1] + s[i];
			
		for(int i=1;i<=n;i++)
			System.out.println(x[i]);




	}

}
