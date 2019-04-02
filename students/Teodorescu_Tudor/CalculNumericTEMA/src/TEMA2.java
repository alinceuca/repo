import java.util.Scanner;
public class TEMA2 {
	
	public static void main(String[] args) {
		Scanner sc  = new Scanner(System.in);
		System.out.println("n=");
		int n = sc.nextInt();
		int a[] = new int[100];
		int b[] = new int[100];
		int c[] = new int [100];
		int d[] = new int [100];
		double x[] = new double [100];
		int r[] = new int[100];
		int s[]=new int[100];
		for(int i=1;i<=n;i++)
			a[i] = sc.nextInt();
		for(int i=2;i<n;i++)
			b[i] = sc.nextInt();
		for(int i=1;i<=n-1;i++)
			c[i] = sc.nextInt();
		for(int i=1;i<=n;i++)
			d[i] = sc.nextInt();
		r[1] = -c[1]/a[1];
		s[1] = d[1]/a[1];
		for(int i=2;i<n-1;i++)
		{
			r[i] = -c[i]/(b[i]*r[i-1]+a[i-1]);
			s[i] = (d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
			
		}
		for(int i=n;i>=2;i--)
			x[i] = (d[i]-b[i]*s[i-1])/(b[i]*r[i-1]+a[i]);
		x[1] = (-c[1]*x[2])/a[1]+d[1]/a[1];

		for(int i=1;i<=n;i++)
			System.out.print(x[i]+" ");
	}

}
