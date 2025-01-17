import 'package:flutter/material.dart';
import 'package:routiner/src/component/routine_card.dart';
import 'package:table_calendar/table_calendar.dart';
import 'package:routiner/src/repo/task_repo.dart';
import 'package:provider/provider.dart';
import 'package:intl/intl.dart';

class HomePage extends StatefulWidget {
  @override
  _HomePage createState() => _HomePage();
}

/// Displays a list of SampleItems.
class _HomePage extends State<HomePage> {
  //static const routeName = '/home';

  DateTime _selectedDay = DateTime.now();
  DateTime _focusedDay = DateTime.now();
  var routineCards;

  @override
  void initState() {
    super.initState();  
    final now = DateTime.now();
    Provider.of<TaskProvider>(context, listen: false).fetchData(now.day, now.month, now.year);
  }

  @override
  void dispose() {
    
    super.dispose();
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
        firstDay: DateTime.utc(2025, 1, 1),
        lastDay: DateTime.utc(2030, 12, 31),
        focusedDay: _focusedDay,

        selectedDayPredicate: (day) {
          return isSameDay(_selectedDay, day);
        },

        onDaySelected: (selectedDay, focusedDay) {
          //if(selectedDay == focusedDay) return;
          setState(() {
            _selectedDay = selectedDay;
            _focusedDay = focusedDay; // update `_focusedDay` here as well
            Provider.of<TaskProvider>(context, listen: false).fetchData(_focusedDay.day, _focusedDay.month, _focusedDay.year);
          });
        },

        /*
        calendarBuilders: CalendarBuilders(
          defaultBuilder: (context, day, focusedDay) {
            return Center(
              child: ColoredBox(color: Colors.black),
            );
          },
        ),
        */
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
      child: Consumer<TaskProvider>(
        builder: (context, provider, child) {
          routineCards = provider.data.map((task) => TaskSmallCard(
            task: task,
            iconSize: 100,
          )).toList();

          // Concatenate static and dynamic items
          final allItems = [...staticItems, ...routineCards];

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
