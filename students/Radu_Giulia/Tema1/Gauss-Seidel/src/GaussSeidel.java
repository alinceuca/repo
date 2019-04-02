import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Scanner;

public class GaussSeidel {

	public static double a[][];
	public static double b[];
	public static double x[],y[];
	public static int ok=0,n;
	public static double nr_max=1000;
	public static int nr_it=0;
	public static double norm;
	public static double eps=Math.pow(10, -4);
	public static Scanner sc;
	
	static void change() {
		for(int i=0;i<n;i++)
			x[i]=y[i];
	}
	
	static double sum1(int linie)
	{
		double s=0.0;
		for(int i=0;i<linie;i++)
			if(i != linie)
				s+=a[linie][i]*y[i];
		return s;
	}
	
	static double sum2(int linie)
	{
		double s=0.0;
		for(int i=linie;i<n;i++)
			if(i != linie)
				s+=a[linie][i]*x[i];
		return s;
	}
	
	static double norma_max()
	{
		double dif=0;
		double max=Math.abs(x[0]-y[0]);
		for(int i=1;i<n;i++) {
			dif=Math.abs(x[i]-y[i]);
			if(dif>max)
				max=dif;
		}
		return max;
	}
	
	static void Gauss_Seidel()
	{
		do {
			
			for(int i=0;i<n;i++)
				y[i]=(b[i]-sum1(i)-sum2(i))/a[i][i];
			norm=norma_max();
			nr_it++;
			change();
			}while(norm>=eps && nr_it<nr_max);
	}
	
	static void read() throws IOException
	{
		sc = new Scanner(new FileReader(new File("date2.in")));
		n=sc.nextInt();
		a=new double[n][n];
		b=new double[n];
		x=new double[n];
		y=new double[n];
		for(int i=0;i<n;i++)
			for(int j=0;j<n;j++)
				a[i][j]=sc.nextDouble();
		for(int i=0;i<n;i++)
			b[i]=sc.nextDouble();
	}
	
	public static void main(String[] args) {
		
		try {
			read();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		Gauss_Seidel();
		for(int i=0;i<n;i++)
			System.out.println(y[i]+ " ");
		if(norm<eps)
			ok=1;
		System.out.println("Ok: "+ ok);
		

	}

}
