package club.tp0t.oj.Graphql.resolvers;

import club.tp0t.oj.Entity.Bulletin;
import club.tp0t.oj.Graphql.types.BulletinResult;
import club.tp0t.oj.Service.BulletinService;
import com.coxautodev.graphql.tools.GraphQLQueryResolver;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
public class BulletinQuery implements GraphQLQueryResolver {
    @Autowired
    private BulletinService bulletinService;

    public BulletinResult allBulletin(){
        List<Bulletin> bulletins = bulletinService.getAllBulletin();
        BulletinResult  bulletinresult = new BulletinResult ("");
        if(bulletins == null) bulletins = new ArrayList<>();
        bulletinresult.setBulletin(bulletins);
        return bulletinresult;
    }
}