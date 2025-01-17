import 'package:flutter/material.dart';
import 'package:routiner/src/component/edit_routine_panel.dart';
import 'package:routiner/src/model/routine.dart';

import 'package:routiner/src/repo/routine_repo.dart';
import 'package:provider/provider.dart';

class EditPage extends StatefulWidget {
  const EditPage({super.key, });

  @override
  State<EditPage> createState() => _EditPage();
}

class _EditPage extends State<EditPage> {

  var _selectedIndex = -1;
  late Future<void> _fetch;

  List<Routine> _filteredItems = [];
  List<Routine> _routines = []; 
  Widget? _detailPanel;

  @override
  void initState() {
    super.initState();
    _fetchData();
  }

  Future<void> _fetchDataFuture() async {
    final routineProvider = Provider.of<RoutineProvider>(context, listen: false);
    await routineProvider.fetchData();
    _routines = routineProvider.data;
    _filteredItems = _routines;
  }

  void _fetchData() {
    _fetch = _fetchDataFuture();
  }

  void _filterItems(String query) {
    setState(() {
      _filteredItems = _routines.where((item) => item.title.toLowerCase().contains(query.toLowerCase())).toList();
    });
  }

  @override
  Widget build(BuildContext context) {

    final theme = Theme.of(context);

    return LayoutBuilder(
      builder: (context, constraints) {
        return Scaffold(
          appBar: AppBar(
            title: const Text("Hello"),
          ),
          body: Row(
            children: [
              Expanded(
                flex: 1,
                child: Column(
                  children: [
                    Align(
                      child: Padding(
                        padding: EdgeInsets.all(8.0),
                        child: ListTile(
                          leading: Icon(Icons.search),
                          title: TextField(
                            onChanged: (query) {
                              _filterItems(query);
                            },
                            decoration: InputDecoration(
                              hintText: 'Search',
                            ),
                          ),
                        ),
                      ),
                    ),
                  
                    Expanded(
                      child: _routinesListPanel(context)
                    ),
                    Align(
                      child: ListTile(
                        contentPadding: EdgeInsets.symmetric(horizontal: 16.0, vertical: 10.0),
                        leading: Icon(Icons.add),
                        title: Text("Add"),
                        onTap: (() {
                          setState(() {
                            _selectedIndex = -1;
                            _detailPanel = NewRoutinePanel(fetchDataTrigger: _fetchData);
                          });
                          _fetchData();
                        }),
                      ),   
                    ),
                  ],
                ),
              ),
              Expanded(
                flex: 3,
                child: Container(
                  child: _detailPanel
                )
              ),
            ],
          ),
        );
      }
    );
  }


  Widget _routinesListPanel(context) {
    
    final theme = Theme.of(context);  
    // final routineProvider = Provider.of<RoutineProvider>(context, listen: false);
    // final routines = routineProvider.data;

    return FutureBuilder<void>(
      future: _fetch,
      builder: (context, snapshot) {
        // if (snapshot.connectionState == ConnectionState.waiting) {
          
        // } else 
        if (snapshot.hasError) {
          return Text('Error: ${snapshot.error}'); 
        } else if (snapshot.connectionState == ConnectionState.done) { 
          // Future completed successfully
          return ListView.builder(
            itemCount: _filteredItems.length,
            itemBuilder: (context, index) {
              return ListTile(
                leading: Icon(IconData(_filteredItems[index].icon, fontFamily: 'MaterialIcons')),
                title: Text(_filteredItems[index].title),
                onTap: () {
                  setState(() {
                    _selectedIndex = index;
                    _detailPanel = NewRoutinePanel(routine: _filteredItems[_selectedIndex], fetchDataTrigger: _fetchData,);
                  });
                },
                selected: index == _selectedIndex,
                selectedTileColor: theme.primaryColor,
                selectedColor: theme.cardColor,
              );
            },
          );
        } else {
          return Text('Unexpected state'); 
        }
      },
    );

  }
}