import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:provider/provider.dart';
import 'package:routiner/src/component/setting.dart';
import 'package:routiner/src/model/task.dart';
import 'package:routiner/src/repo/task_repo.dart';

class TaskSmallCard extends StatefulWidget {

  final Task task;
  final double iconSize;

  const TaskSmallCard({ 
    Key? key, 
    required this.task,
    required this.iconSize,

  }): super(key: key);

  @override
  State<TaskSmallCard> createState() => _TaskSmallCard();
}

/// Displays a list of SampleItems.
class _TaskSmallCard extends State<TaskSmallCard> {

  bool? status = false;

  @override
  void initState() {
    super.initState();  
    setState(() {
      status = widget.task.status;
    });
  }

  @override
  void didUpdateWidget(covariant TaskSmallCard oldWidget) {
    super.didUpdateWidget(oldWidget);
    setState(() {
      status = widget.task.status;
    });
  }

  @override
  Widget build(BuildContext context) {

    //print()

    final theme = Theme.of(context);  
    final titleStyle = theme.textTheme.titleLarge!.copyWith(
      color: Colors.black,
    );

    return Card(
        child: Padding(
          padding: EdgeInsets.all(10),
          child: 
            Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                Expanded(
                  flex: 4,
                  child: Row(
                    children: [
                      Icon(
                        IconData(widget.task.routine?.icon ?? 50000, fontFamily: 'MaterialIcons'),
                        size: widget.iconSize,
                      ),
                      SizedBox(width: 30,),
                      Column(
                        mainAxisAlignment: MainAxisAlignment.start,
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            widget.task.routine?.title ?? "No title",
                            style: titleStyle,
                          ),
                          Text(
                            widget.task.routine?.explain ?? '',
                          ),
                          
                          if(widget.task.begin != null && widget.task.due != null) 
                            Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children : [
                              Text(
                                  "Start: ${DateFormat(Setting.dateFormat).format(widget.task.begin!)}"
                                ),
                                Text(
                                  "Due: ${DateFormat(Setting.dateFormat).format(widget.task.due!)}"
                                )
                              ],
                            )
                          else Text("Begin or Due is broken")
                          
                          
                        ]
                      ),
                    ],

                  ),
                ),
                
                Expanded(
                  flex : 1,
                  child: Consumer<TaskProvider>(
                    builder: (context, provider, child) {
                      return Checkbox(
                        value: status,
                        onChanged: (bool? value) {
                          setState(() {
                            status = value;
                          });
                          Task task = Task(id: widget.task.id, status: status!);
                          Provider.of<TaskProvider>(context, listen: false).checkTask(widget.task.id, task);
                        },
                      );
                    }
                  )
                ),
              ],
            ),
        ), 
      );
  }

}
