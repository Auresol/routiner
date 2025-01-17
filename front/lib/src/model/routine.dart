import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

class Routine {
  int? id;
  final String title;
  String? explain;
  final int icon;
  DateTime activeDate = DateTime(2000, 1, 1);
  int dueIn = 1;
  bool forceReset = false;
  final RoutineMode routineMode;
  int? dayInWeekly;
  int? frequency;
  bool resetOnMonth = false;

  DateTime? createdAt = DateTime(2000, 1, 1);
  DateTime? updateAt = DateTime(2000, 1, 1);

  Routine({
    this.id,
    required this.title,
    this.explain,
    required this.icon,
    required this.activeDate,
    required this.dueIn,
    this.forceReset = false,
    required this.routineMode,
    this.dayInWeekly,
    this.frequency,
    this.resetOnMonth = false,
    this.createdAt,
    this.updateAt 
  });

  Routine.empty({
    this.title = "empty",
    this.icon = 50000,
    this.routineMode = RoutineMode.NULL,
  });

  Map<String, dynamic> toJson() {
    return {
      'id' : id,
      'title': title,
      'explain': explain,
      'icon_code_point': icon,
      //DateTime(2000,1,1).toIso8601String()
      'active_date': DateFormat("yyyy-MM-ddTHH:mm:ss'Z'").format(activeDate),
      'due_in': dueIn,
      'force_reset': forceReset,
      'routine_mode': routineMode.index,
      'day_in_weekly': dayInWeekly,
      'frequency': frequency,
      'reset_on_month': resetOnMonth
    };
  }

  factory Routine.fromJson(Map<String, dynamic> parsedJson){
    return Routine(
      id: parsedJson['id'],
      title: parsedJson['title'],
      explain: parsedJson['explain'],
      icon: parsedJson['icon_code_point'],
      activeDate: DateTime.parse(parsedJson['active_date']).toLocal(),
      dueIn: parsedJson['due_in'],
      forceReset: parsedJson['force_reset'],
      routineMode: RoutineMode.values[parsedJson['routine_mode']],
      dayInWeekly: parsedJson['day_in_weekly'],
      frequency: parsedJson['frequency'],
      resetOnMonth: parsedJson['reset_on_month'],
      createdAt: DateTime.parse(parsedJson['created_at']).toLocal(),
      updateAt: DateTime.parse(parsedJson['updated_at']).toLocal(),
    );
  }

}

enum RoutineMode { NULL, WEEKLY, PERIOD, TODO }

// final routineCardTestData = 
// [
//   Routine(id: 1, title: "Running", icon: Icons.directions_run_rounded.codePoint, explain: "For 30 mins"),
//   Routine(id: 2, title: "Workout", icon: Icons.accessible_forward_rounded.codePoint, explain: "12 sets"),
//   Routine(id: 3, title: "Reading", icon: Icons.book.codePoint, explain: "3 books"),
//   Routine(id: 4, title: "Calculus", icon: Icons.calculate.codePoint, explain: "1 chapter"),
// ];

