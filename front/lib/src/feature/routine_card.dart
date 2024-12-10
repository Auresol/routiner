import 'package:flutter/material.dart';

class RoutineSmallCard extends StatefulWidget {

  final String title;
  final IconData icon;
  final double iconSize;
  final String? explain;

  const RoutineSmallCard({ 
    Key? key, 
    required this.title, 
    required this.icon, 
    required this.iconSize,
    this.explain
  }): super(key: key);

  @override
  State<RoutineSmallCard> createState() => _RoutineSmallCard();
}

/// Displays a list of SampleItems.
class _RoutineSmallCard extends State<RoutineSmallCard> {

  bool? isChecked = false;
  /*
  _RoutineSmallCard({
    super.key,
    required this.title,
    required this.icon,
    required this.explain,
  });

  final String title;
  final IconData icon;
  final String explain;
  */

  @override
  Widget build(BuildContext context) {

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
                        widget.icon,
                        size: widget.iconSize,
                      ),
                      SizedBox(width: 30,),
                      Column(
                        mainAxisAlignment: MainAxisAlignment.start,
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text(
                            widget.title,
                            style: titleStyle,
                          ),
                          Text(
                            widget.explain ?? '',
                          ),
                        ],
                      ),
                    ],

                  ),
                ),
                
                Expanded(
                  flex : 1,
                  child: 
                    Checkbox(
                      value: isChecked,
                      onChanged: (bool? value) {
                        setState(() {
                          isChecked = value;
                        });
                      },
                    ),
                ),
              ],
            ),
        ), 
      );
  }
}
