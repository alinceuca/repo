import java.util.Scanner;
public class SistemTridiagonalIterativ {

	public static void tridiagonalIt(int n,double a[],double b[],double c[],double d[]) {

		double x[] = new double [n+1];
		double r[] = new double [n];
		double s[] = new double [n];


		if(a[1]!=0)
		{
			r[1]=-c[1]/a[1];
			s[1]=d[1]/a[1];
		}

		for(int i=1;i<=n-2;i++) {
			r[i + 1] = -c[i + 1] / (b[i] * r[i] + a[i + 1]);
			s[i + 1] = (d[i + 1] - b[i] * s[i]) / (b[i] * r[i] + a[i + 1]);
		}

		x[n] = (d[n]-b[n-1]*s[n-1])/(b[n-1]*r[n-1]+a[n]);

		for(int i=n-1;i>=1;i--) {
			x[i] = r[i]*x[i+1]+s[i];
		}

		System.out.println("Solutiile sunt: ");

		for(int i=1;i<=n;i++) {
			System.out.println("x[" + i + "]: " + (double) Math.round(x[i] * 100000d) / 100000d);
		}

	}

	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);

		double a[],b[],c[],d[];
		int n=sc.nextInt();
		a = new double[n+1];
		b = new double[n];
		c = new double[n];
		d = new double[n+1];

		for(int i=1;i<=n;i++)
			a[i]=sc.nextDouble();

		for(int i=1;i<=n-1;i++)
			b[i]=sc.nextDouble();

		for(int i=1;i<=n-1;i++)
			c[i]=sc.nextDouble();

		for(int i=1;i<=n;i++)
			d[i]=sc.nextDouble();

		tridiagonalIt(n, a, b, c, d);
	}

}