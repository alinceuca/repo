import java.util.Scanner;
public class MetGaussSeidel {
	public static final int maxIt_CONST=100;
	public static double x[],b[],a[][],y[];
	
	public static void main(String[]args) {
		Scanner sc=new Scanner(System.in);
		int n=sc.nextInt(),nrIt=0,OK=0;
		double norm,s,eps=0.001,s1;
		a=new double[n][n];
		b=new double[n];
		x=new double[n];
		y=new double[n];
		for(int i=0;i<n;i++)
			for(int j=0;j<n;j++)
				a[i][j]=sc.nextDouble();
		for(int i=0;i<n;i++)
		{
			b[i]=sc.nextDouble();
			x[i]=0;
		}
		 do{
			for(int i=0;i<n;i++)
			{
				s=0;
				s1=0;
				for(int j=0;j<i;j++)
					s=s+a[i][j] * x[j];
				for(int j=i;j<n;j++)
					s1=s1+a[i][j]*x[j];
				y[i]=(b[i]-s-s1)/a[i][i];
				x[i]=y[i];
			}
			norm=0;
			for (int i=0;i<n;i++)
				norm = norm + Math.abs(y[i] - x[i]);
			nrIt++;
		}while(norm>=eps && nrIt<maxIt_CONST);
		if(norm<eps)
			OK=1;
		if(OK==1)
			System.out.println("Nu s-a atins numarul maxim de iteratii.");
		else System.out.println("S-a atins numarul maxim de iteratii.");
		System.out.println("Solutia aproximativa este ");
		for(int i=0;i<n;i++)
			System.out.println("y["+i+"]="+y[i]);
		}

}
