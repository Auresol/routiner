import 'dart:io';

import 'package:flutter/material.dart';
import 'package:routiner/src/model/block.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

class BlockProvider extends ChangeNotifier {

  static final url = 'http://localhost:8080/api/block';

  // final queryParam = {"test" : "123"};
  // static final baseUrl = Uri.http(host, path, queryParam);

  List<Block> _data = [];

  List<Block> get data => _data;

  Future<void> fetchData() async {
    // Fetch data from your local server
  try {
    final response = await http.get(Uri.parse("${url}s"));
    if (response.statusCode == 200) {
      final List<dynamic> data = jsonDecode(response.body);
      _data = data.map((result) {
        try {
          return Block(
            id: result['id'],
            title: result['title'],
            explain: result['explain'],
            icon: result['icon_code_point'],
          );
        } catch (e) {
          print('Error parsing item: $e');
          return Block(title: 'Error', explain: 'Error parsing data', icon: Icons.error.codePoint);
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

  Future<void> createBlock(Block block) async {
    try {
      final body = jsonEncode(block.toJson());

      final response = await http.post(
        Uri.parse(url),
        headers: {'Content-Type': 'application/json'},
        body: body,
      );

      if (response.statusCode == 201)  {
        // Successful creation
        print('Block created successfully');
      } else {
        // Handle error
        print('Error creating block: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }

  Future<void> deleteBlock(int block_id) async {

    try {
      final response = await http.delete(
        Uri.parse("$url/$block_id"),
        headers: {'Content-Type': 'application/json'},
      );
      
      if (response.statusCode == 200)  
        {
        // Successful creation
        print('Block deleted successfully');
      } else {
        // Handle error
        print('Error deleting block: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }

  Future<void> revertBlock(int block_id) async {

    try {
      final response = await http.post(
        Uri.parse("$url/revert/$block_id"),
        headers: {'Content-Type': 'application/json'},
      );
      
      if (response.statusCode == 200)  
        {
        // Successful creation
        print('Block reverted successfully');
      } else {
        // Handle error
        print('Error reverted block: ${response.statusCode}');
      }
    } catch (e) {
      print("Error: Unexpected error: $e");
    }
  }
}