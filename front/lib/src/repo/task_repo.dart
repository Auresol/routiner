import 'dart:io';

import 'package:flutter/material.dart';
import 'package:routiner/src/model/routine.dart';
import 'package:routiner/src/model/task.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class TaskProvider extends ChangeNotifier {

  static final url = 'http://localhost:8080/api/task';

  // final queryParam = {"test" : "123"};
  // static final baseUrl = Uri.http(host, path, queryParam);

  List<Task> _data = [];

  List<Task> get data => _data;

  Future<void> fetchData(int d, int m, int y) async {
    // Fetch data from your local server
  try {
    final response = await http.get(Uri.parse("${url}/date?d=${d}&y=${y}&m=${m}"));
    if (response.statusCode == 200) {
      
      final List<dynamic> data = jsonDecode(response.body);
      _data = data.map((result) {
        try {
          return Task.fromJson(result);
        } catch (e) {
          print('Error parsing task: $e');
          return Task.empty();
        }
        }).toList();
        notifyListeners();
      } else {
        print("Error fetching : Status code ${response.statusCode}");
        // Handle server-side error
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
      // Handle other unexpected errors
    } finally {
      // Code that always executes, regardless of exceptions
    }
  }

  Future<void> checkTask(int task_id, Task task) async {

    final body = jsonEncode(task.toJson());

    try {
      final response = await http.put(
        Uri.parse("$url/${task_id}"),
        headers: {'Content-Type': 'application/json'},
        body: body,
      );
      
      if (response.statusCode == 200) {
        // Successful creation
        print('Task updated successfully');
      } else {
        // Handle error
        print('Error reverted routine: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }
}