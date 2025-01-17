import 'dart:io';

import 'package:flutter/material.dart';
import 'package:routiner/src/model/routine.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class RoutineProvider extends ChangeNotifier {

  static final url = 'http://localhost:8080/api/routine';

  // final queryParam = {"test" : "123"};
  // static final baseUrl = Uri.http(host, path, queryParam);

  List<Routine> _data = [];
  List<Routine> get data => _data;

  Future<void> fetchData() async {
    // Fetch data from your local server
  try {
    final response = await http.get(Uri.parse("${url}s"));
    if (response.statusCode == 200) {
      final List<dynamic> data = jsonDecode(response.body);
      _data = data.map((result) {
        try {
          return Routine.fromJson(result);
        } catch (e) {
          print('Error parsing routine: $e');
          return Routine.empty();
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

  Future<void> createRoutine(Routine routine) async {
    try {
      final body = jsonEncode(routine.toJson());

      final response = await http.post(
        Uri.parse(url),
        headers: {'Content-Type': 'application/json'},
        body: body,
      );

      if (response.statusCode == 201)  {
        // Successful creation
        print('Routine created successfully');
      } else {
        // Handle error
        print('Error creating routine: ${response.statusCode}');
        print('Return:  ${response.body}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }

  Future<void> deleteRoutine(int routine_id) async {

    try {
      final response = await http.delete(
        Uri.parse("$url/$routine_id"),
        headers: {'Content-Type': 'application/json'},
      );
      
      if (response.statusCode == 200)  
        {
        // Successful creation
        print('Routine deleted successfully');
      } else {
        // Handle error
        print('Error deleting routine: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }

  Future<void> revertRoutine(int routine_id) async {

    try {
      final response = await http.post(
        Uri.parse("$url/revert/$routine_id"),
        headers: {'Content-Type': 'application/json'},
      );
      
      if (response.statusCode == 200) {
        // Successful creation
        print('Routine reverted successfully');
      } else {
        // Handle error
        print('Error reverted routine: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }

  Future<void> checkTask(int routine_id) async {

    try {
      final response = await http.post(
        Uri.parse("$url/revert/$routine_id"),
        headers: {'Content-Type': 'application/json'},
      );
      
      if (response.statusCode == 200) {
        // Successful creation
        print('Routine reverted successfully');
      } else {
        // Handle error
        print('Error reverted routine: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }
}