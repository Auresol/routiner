import 'package:flutter/material.dart';
import 'package:routiner/src/component/detail_block_panel.dart';
import 'package:routiner/src/component/new_block_panel.dart';
import 'package:routiner/src/model/block.dart';

import 'package:routiner/src/repo/block_repo.dart';
import 'package:provider/provider.dart';

class EditPage extends StatefulWidget {
  const EditPage({super.key, });

  @override
  State<EditPage> createState() => _EditPage();
}

class _EditPage extends State<EditPage> {

  var _selectedIndex = 0;
  List<Block> _filteredItems = [];
  Widget _detailPanel = const NewBlockPanel();

  @override
  void initState() {
    super.initState();
  }

  void _filterItems(List<Block> allBlock, String query) {
    setState(() {
      _filteredItems = allBlock.where((item) => item.title.toLowerCase().contains(query.toLowerCase())).toList();
    });
  }

  @override
  Widget build(BuildContext context) {

    final theme = Theme.of(context);  

    final blockProvider = Provider.of<BlockProvider>(context); // Access Provider

    // Get blocks from provider on build (assuming one time fetch)
    final List<Block> _blocks = blockProvider.data ?? [];
    _filteredItems = _blocks;

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
                              _filterItems(_blocks, query);
                            },
                            decoration: InputDecoration(
                              hintText: 'Search',
                            ),
                          ),
                        ),
                      ),
                    ),
                  
                    Expanded(
                      child: ListView.builder(
                          itemCount: _filteredItems.length,
                          itemBuilder: (context, index) {
                            return ListTile(
                              leading: Icon(IconData(_filteredItems[index].icon, fontFamily: 'MaterialIcons')),
                              title: Text(_filteredItems[index].title),
                              onTap: () {
                                setState(() {
                                  _selectedIndex = index;
                                  _detailPanel = DetailBlockPanel(block: _filteredItems[_selectedIndex]);
                                });
                              },
                              selected: index == _selectedIndex,
                              selectedTileColor: theme.primaryColor,
                              selectedColor: theme.cardColor,
                            );
                          },
                        ),
                    ),
                    Align(
                      child: ListTile(
                        contentPadding: EdgeInsets.symmetric(horizontal: 16.0, vertical: 10.0),
                        leading: Icon(Icons.add),
                        title: Text("Add"),
                        onTap: (() {
                          setState(() {
                            _detailPanel = const NewBlockPanel();
                          });
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
}