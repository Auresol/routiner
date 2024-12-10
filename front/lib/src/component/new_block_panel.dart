import 'package:flutter/material.dart';
import 'package:flutter_iconpicker/Models/configuration.dart';
import 'package:routiner/src/model/block.dart';
import 'package:flutter_iconpicker/flutter_iconpicker.dart';

import 'package:routiner/src/repo/block_repo.dart';
import 'package:provider/provider.dart';


class NewBlockPanel extends StatefulWidget {
    const NewBlockPanel({super.key, });

    @override
    _NewBlockPanel createState() => _NewBlockPanel();
}

/// Displays a list of SampleItems.
class _NewBlockPanel extends State<NewBlockPanel> {
  //static const routeName = '/home';

  static Block? _currentData;

  @override
  void initState() {
    super.initState();  
  }

  final _formKey = GlobalKey<FormState>();
  final _titleController = TextEditingController();
  final _explainController = TextEditingController();
  final _iconController =  TextEditingController();

  IconData _iconData = Icons.book;

  _pickIcon() async {
    IconPickerIcon? icon = await showIconPicker(
        context,
        configuration: SinglePickerConfiguration(
          iconPackModes: [IconPack.roundedMaterial],
        ),
    );
    
    if(icon?.data != null) {
      _iconData = icon!.data;
    }

    setState(() {});

    debugPrint('Picked Icon:  $icon');
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Form(
        key: _formKey,
        child: Column(
          children: [
            
            Expanded(
              child: Column(
                children: [
                  
                  Align(
                    child: Row(
                      children: [
                        Expanded(
                          flex: 2,
                          child: Column(
                            children: [
                              Icon(
                                _iconData,
                              ),
                              ElevatedButton(
                                onPressed: _pickIcon,
                                child: const Text('Open IconPicker'),
                              ),
                            ],
                          ),
                        ),
                        Expanded(
                          flex: 5,
                          child: Column(
                            children: [
                              TextFormField(
                                controller: _titleController,
                                decoration:  InputDecoration(labelText: 'Title'),
                                validator: (value) {
                                  if (value == null || value.isEmpty) {
                                    return 'Please enter a title';
                                  }
                                  return  null;
                                },
                              ),
                              TextFormField(
                                controller: _explainController,
                                decoration:  InputDecoration(labelText: 'Explain'),
                              ),
                          
                            ],
                          ),
                        ),
                      ],
                    ),
                  ),
                ],
              ),
            ),
      
            Consumer<BlockProvider>(
              builder: (context, provider, child) {
                return ElevatedButton(
                    onPressed: () {
                      if (_formKey.currentState!.validate()) {
                        // Process the form data, create a new Block object, and close the panel
                        // Navigator.pop(context, Block(
                        //   title: _titleController.text,
                        //   explain: _explainController.text,
                        //   icon: _iconData.codePoint,
                        // ));
                        Block newBlock = Block(
                          title: _titleController.text,
                          explain: _explainController.text,
                          icon: _iconData.codePoint,
                        );

                        provider.createBlock(newBlock);
                      }
                    },
                    child: Text('Create'),
                );
              }
            )

          ],
        ),
      ),
    );
  }
}
