package com.dao;
import com.beans.*;
import com.JDBC.*;
import java.sql.*;

public class pruduce_dao extends abstruct_dao {
    produce Produce;

    public pruduce_dao(produce inputproduce){
        super();
        this.Produce=inputproduce;
    }

    public pruduce_dao(Connection conn,produce inputproduce){
        super(conn);
        this.Produce=inputproduce;
    }

}
