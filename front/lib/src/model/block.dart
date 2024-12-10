import 'package:flutter/material.dart';

class Block {
  int? id;
  final String title;
  String? explain;
  final int icon;

  Block({
    this.id,
    required this.title,
    required this.icon,
    this.explain,
  });

  Map<String, dynamic> toJson() {
    return {
      'id' : id,
      'title': title,
      'explain': explain,
      'icon_code_point': icon,
    };
  }
}

final routineCardTestData = 
[
  Block(id: 1, title: "Running", icon: Icons.directions_run_rounded.codePoint, explain: "For 30 mins"),
  Block(id: 2, title: "Workout", icon: Icons.accessible_forward_rounded.codePoint, explain: "12 sets"),
  Block(id: 3, title: "Reading", icon: Icons.book.codePoint, explain: "3 books"),
  Block(id: 4, title: "Calculus", icon: Icons.calculate.codePoint, explain: "1 chapter"),
];

