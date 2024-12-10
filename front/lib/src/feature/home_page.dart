import 'package:flutter/material.dart';
import 'package:routiner/src/feature/routine_card.dart';
import 'package:table_calendar/table_calendar.dart';
import 'package:routiner/src/repo/block_repo.dart';
import 'package:provider/provider.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePage createState() => _HomePage();
}

/// Displays a list of SampleItems.
class _HomePage extends State<HomePage> {
  //static const routeName = '/home';

  @override
  void initState() {
    super.initState();  

    Provider.of<BlockProvider>(context, listen: false).fetchData();
  }

  @override
  Widget build(BuildContext context) {

    final theme = Theme.of(context);  
    final welcomeTextStyle = theme.textTheme.displayMedium!.copyWith(
      color: Colors.black,
    );
    final titleTextStyle = theme.textTheme.titleMedium!.copyWith(
      color: Colors.black,
    );

    List<Widget> staticItems = [
      Padding(
        padding: const EdgeInsets.all(20.0),
        child: Text( 
          "Welcome",
          style: welcomeTextStyle,
        ),
      ),

      TableCalendar(
        firstDay: DateTime.utc(2010, 10, 16),
        lastDay: DateTime.utc(2030, 3, 14),
        focusedDay: DateTime.now(),
      ),

      Padding(
        padding: const EdgeInsets.all(20.0),
        child: Text( 
          "Today's task",
          style: titleTextStyle,
        ),
      ),
      SizedBox(height: 3.0),
    ];

    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Consumer<BlockProvider>(
        builder: (context, provider, child) {
          final dynamicItems = provider.data.map((block) => RoutineSmallCard(
            title: block.title,
            icon: IconData(block.icon, fontFamily: 'MaterialIcons'), // Assuming your Block model has an icon property
            iconSize: 100,
            explain: block.explain,
          )).toList();

          // Concatenate static and dynamic items
          final allItems = [...staticItems, ...dynamicItems];

          return ListView.builder(
            itemCount: allItems.length,
            itemBuilder: (context, index) {
              return allItems[index];
            },
          );
        }
      ),
    );
  }
}
