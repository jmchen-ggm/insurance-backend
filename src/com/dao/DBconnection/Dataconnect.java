package com.dao.DBconnection;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.Statement;

public class Dataconnect {


    private static final String driver = "org.sqlite.JDBC";; //数据库驱动
    //连接数据库的URL地址
    private static final String myurl = "jdbc:sqlite:insuranceDB.db";

    private static Connection conn=null;

    static
    {
        try
        {
            Class.forName(driver);

        }
        catch (Exception ex)
        {
            ex.printStackTrace();
            System.out.println("加载驱动失败");
        }

    }

    public static Connection getConnection() throws Exception
    {

        if(conn==null)
        {
            conn = DriverManager.getConnection(myurl);
        }
        return conn;
    }

    public static void main( String args[] )
    {
        Connection c = null;
        Statement stmt = null;
        try {
            c=getConnection();
            System.out.println("Opened database successfully");
            stmt = c.createStatement();
            // String sql = "DELETE from produce where rowid=2;";
            //  stmt.executeUpdate(sql);
            //  c.commit();

            ResultSet rs = stmt.executeQuery( "SELECT * FROM produce;" );
            while ( rs.next() ) {
                int id = rs.getInt("id");
                String  name = rs.getString("name");
                int age  = rs.getInt("rank");
                int  address = rs.getInt("money");
                float salary = rs.getFloat("score");
                float salary1 = rs.getFloat("lpercenta");
                float salary2 = rs.getFloat("lpercentb");
                float salary3 = rs.getFloat("skind");
                float salary4 = rs.getFloat("sgroup");
                float salary5 = rs.getFloat("snumber");
                float salary6 = rs.getFloat("lgroup");
                float salary7 = rs.getFloat("lnumber");
                String salary8 = rs.getString("addition");
                System.out.println( "rowid = " + id );
                System.out.println( "name = " + name );
                System.out.println( "rank = " + age );
                System.out.println( "money = " + address );
                System.out.println( "score = " + salary );
                System.out.println( "score = " + salary1 );
                System.out.println( "score = " + salary2 );
                System.out.println( "score = " + salary3 );
                System.out.println( "score = " + salary4 );
                System.out.println( "score = " + salary5 );
                System.out.println( "score = " + salary6 );
                System.out.println( "score = " + salary7 );
                System.out.println( "score = " + salary8 );

                System.out.println();
            }
            rs.close();
            stmt.close();
            c.close();
        } catch ( Exception e ) {
            System.err.println( e.getClass().getName() + ": " + e.getMessage() );
            System.exit(0);
        }
        System.out.println("Operation done successfully");
        //close(c,stmt);
    }
}
