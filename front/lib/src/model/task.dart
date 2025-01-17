import 'package:routiner/src/model/routine.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';

class Task {
  final int id;
	DateTime? begin;
	DateTime? due;
	String? detail;
	bool status = false;   
  int? routineId;
  Routine? routine;

  Task.empty({
    this.id = 0,
    this.routineId = 0,
    this.status = false,
  });

  Task({
    required this.id,
    this.routineId,
    this.begin,
    this.due,
    this.detail,
    required this.status,
    this.routine,
  });

  Map<String, dynamic> toJson() {
    return {
      'id' : id,
      'routine_id': routineId,
      'begin': begin,
      'due': due,
      'detail' : detail,
      'status' : status,
      'routine' : routine?.toJson()
    };
  }

  factory Task.fromJson(Map<String, dynamic> parsedJson){
    return Task(
      id: parsedJson['id'],
      routineId: parsedJson['routine_id'],
      begin: DateTime.parse(parsedJson['begin']).toLocal(),
      due: DateTime.parse(parsedJson['due']).toLocal(),
      detail: parsedJson['detail'],
      status: parsedJson['status'],
      routine: Routine.fromJson(parsedJson['routine'])
    );
  }
}
