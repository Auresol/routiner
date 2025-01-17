import 'dart:ffi';

import 'package:flutter/material.dart';
import 'package:flutter_iconpicker/Models/configuration.dart';
import 'package:intl/intl.dart';
import 'package:routiner/src/model/routine.dart';
import 'package:flutter_iconpicker/flutter_iconpicker.dart';

import 'package:routiner/src/repo/routine_repo.dart';
import 'package:provider/provider.dart';

import 'package:date_picker_plus/date_picker_plus.dart';
import 'package:routiner/src/component/setting.dart';

class NewRoutinePanel extends StatefulWidget {

    final void Function() fetchDataTrigger;
    final Routine? routine;

    const NewRoutinePanel({super.key, this.routine, required this.fetchDataTrigger});

    @override
    _NewRoutinePanel createState() => _NewRoutinePanel();
}

/// Displays a list of SampleItems.
class _NewRoutinePanel extends State<NewRoutinePanel> {
  //static const routeName = '/home';

  @override
  void initState() {
    super.initState();  
    updateAllFormField();
  }

  @override
  void didUpdateWidget(covariant NewRoutinePanel oldWidget) {
    super.didUpdateWidget(oldWidget);
    updateAllFormField();
  }

  final _formKey = GlobalKey<FormState>();
  final _titleController = TextEditingController();
  final _explainController = TextEditingController();
  final _iconController =  TextEditingController();
  final _dueinController = TextEditingController();
  var _activeDate = DateTime(2000, 1, 1);
  final _activeDateController = TextEditingController();
  final _frequencyController = TextEditingController();
  RoutineMode? _routineMode = RoutineMode.TODO;
  var _forceResetCheckbox = false;
  var _resetOnMonthCheckbox = false;
  var _isCreatingNewRoutine = true;
  var isDeleted = false;

  void updateAllFormField() {
      _titleController.text = widget.routine?.title ?? '';
      _explainController.text = widget.routine?.explain ?? '';
      _iconController.text = widget.routine?.icon.toString() ?? '';
      _dueinController.text = widget.routine?.dueIn.toString() ?? '1';
      _activeDate = widget.routine?.activeDate.toLocal() ?? DateTime.now();
      _activeDateController.text = DateFormat(Setting.dateFormat).format(_activeDate);
      _forceResetCheckbox = widget.routine?.forceReset ?? false;
      _resetOnMonthCheckbox = widget.routine?.resetOnMonth ?? false;
      _routineMode = widget.routine?.routineMode ?? RoutineMode.TODO;
      _selectedDaysNumber = widget.routine?.dayInWeekly ?? 0;

      for(int i = 0;i < 7;i++){
        _selectedDays[i] = ((_selectedDaysNumber >> i) & 1) == 1;
      }
      
      _frequencyController.text = widget.routine?.frequency.toString() ?? '1';
      _isCreatingNewRoutine = widget.routine == null;

      if(widget.routine != null) {
        _formKey.currentState?.validate();
      }
  }

  List<bool> _selectedDays = List.filled(7, false); 
  var _selectedDaysNumber = 0;
  void _handleDaySelected(int index) {
    setState(() {
      _selectedDaysNumber = _selectedDaysNumber ^ (1 << index);
    });
    //print(_selectedDaysNumber);
  }

  final daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];

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
  }

  void pickDate() async {
    final date = await showDatePickerDialog(
      context: context,
      minDate: DateTime(2025, 1, 1),
      maxDate: DateTime(2026, 12, 31),
      width: 450,
      height: 450,
    );

    if(date == null) return;
    
    _activeDate = date;
    _activeDateController.text = DateFormat(Setting.dateFormat).format(_activeDate);
  }
  @override
  Widget build(BuildContext context) {

    final theme = Theme.of(context);  
    final labelTextStyle = theme.textTheme.bodyLarge!;

    return Padding(
      padding: const EdgeInsets.all(32.0),
      child: Form(
        key: _formKey,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Column(
              children: [
                Row(
                  children: [
                    Expanded(
                      flex: 2,
                      child: Column(
                        children: [
                          Icon(
                            _iconData,
                            size: 50,
                          ),
                          ElevatedButton(
                            onPressed: _pickIcon,
                            child: const Text('IconPicker'),
                          ),
                        ],
                      ),
                    ),
                    Expanded(
                      flex: 6,
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
                        ],
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 10,),
                Column(
                  children: [
                    TextFormField(
                      controller: _explainController,
                      decoration:  InputDecoration(labelText: 'Explain'),
                    ),
                    SizedBox(height: 20,),

                    Row(
                      children: [
                        Expanded(
                          flex: 1,
                          child:  TextFormField(
                            controller: _dueinController,
                            decoration:  InputDecoration(labelText: 'Due in (days)'),
                          ),
                        ),
                        SizedBox(width: 10),
                        Expanded(
                          flex: 2,
                          child:  TextFormField(
                            controller: _activeDateController,
                            decoration:  InputDecoration(labelText: 'Active date'),
                          ),
                        ),
                        SizedBox(width: 10),
                        Expanded(
                          flex : 1,
                          child: ElevatedButton(
                            onPressed: () => pickDate(), 
                            child: Text("Date picker")
                          ),
                        ),
                        SizedBox(width: 10),
                        Expanded(
                          flex: 1,
                          child: Row(
                            children: [
                              Text('Force reset: ', style: labelTextStyle,),
                              Checkbox(
                                value: _forceResetCheckbox,
                                onChanged: (bool? value) {
                                  setState(() {
                                    _forceResetCheckbox = value ?? false;
                                  });
                                },
                              ),
                            ]
                          ),
                        ),
                      ],
                    ),
                    SizedBox(height: 20,),

                    Row(
                      mainAxisAlignment: MainAxisAlignment.start,
                      children: <Widget>[
                        Text(
                          "Routine Mode: ",
                          style: labelTextStyle,
                        ),
                        Expanded(
                          child: ListTile(
                            title: const Text('Weekly'),
                            leading: Radio<RoutineMode>(
                              value: RoutineMode.WEEKLY,
                              groupValue: _routineMode,
                              onChanged: (RoutineMode? value) {
                                setState(() {
                                  _routineMode = value;
                                });
                              },
                            ),
                          ),
                        ),
                        Expanded(
                          child: ListTile(
                            title: const Text('Period'),
                            leading: Radio<RoutineMode>(
                              value: RoutineMode.PERIOD,
                              groupValue: _routineMode,
                              onChanged: (RoutineMode? value) {
                                setState(() {
                                  _routineMode = value;
                                });
                              },
                            ),
                          ),
                        ),
                        Expanded(
                          child: ListTile(
                            title: const Text('TODO'),
                            leading: Radio<RoutineMode>(
                              value: RoutineMode.TODO,
                              groupValue: _routineMode,
                              onChanged: (RoutineMode? value) {
                                setState(() {
                                  _routineMode = value;
                                });
                              },
                            ),
                          ),
                        ),
                      ],
                    ),

                    SizedBox(height: 15,),

                    if (_routineMode == RoutineMode.WEEKLY) 
                      Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [ 
                          Text(
                            "Days in weekly mode ",
                            style: theme.textTheme.titleSmall!,
                          ),
                          SizedBox(height: 20,),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            children: List.generate(7, (index) {
                              return GestureDetector(
                                onTap: () {
                                  setState(() {
                                    _selectedDays[index] = !_selectedDays[index]; 
                                    _handleDaySelected(index); 
                                  });
                                },
                                child: Container(
                                  width: 70, 
                                  height: 40,
                                  decoration: BoxDecoration(
                                    color: _selectedDays[index] ? theme.primaryColor : theme.dialogBackgroundColor, 
                                    borderRadius: BorderRadius.circular(5),
                                    border: Border.all(width: 0.5),
                                  ),
                                  child: Center(
                                    child: Text(
                                      daysOfWeek[index],
                                      style: TextStyle(
                                        color: _selectedDays[index] ? Colors.white : Colors.black,
                                      ),
                                    ),
                                  ),
                                ),
                              );
                            }),
                          ),
                        ],
                      ),

                      if (_routineMode == RoutineMode.PERIOD)
                        Row(
                          children: [
                            Expanded(
                              flex: 2,
                              child:  TextFormField(
                                controller: _frequencyController,
                                decoration:  InputDecoration(labelText: 'Frequency'),
                              ),
                            ),
                            Expanded(
                              flex: 2,
                              child: Row(
                                children: [
                                  Text('Reset at the start of the month: ', style: labelTextStyle,),
                                  Checkbox(
                                  value: _resetOnMonthCheckbox,
                                  onChanged: (bool? value) {
                                    setState(() {
                                      _resetOnMonthCheckbox = value ?? false;
                                    });
                                  },
                                ),
                                ],
                              ),
                            ),
                          ],
                        ),
                      
                  ],
                )
              ],
            ),
      
            Consumer<RoutineProvider>(
              builder: (context, provider, child) {
                if(_isCreatingNewRoutine) {
                  return ElevatedButton(
                      onPressed: () {
                        if (_formKey.currentState?.validate() ?? false) {
                          // Process the form data, create a new Routine object, and close the panel
                          // Navigator.pop(context, Routine(
                          //   title: _titleController.text,
                          //   explain: _explainController.text,
                          //   icon: _iconData.codePoint,
                          // ));
                          if(_isCreatingNewRoutine) {
                        
                            Routine newRoutine = Routine(
                              title: _titleController.text,
                              explain: _explainController.text,
                              icon: _iconData.codePoint,
                              activeDate: _activeDate,
                              dueIn: int.parse(_dueinController.text),
                              forceReset: _forceResetCheckbox,
                              routineMode: _routineMode!,
                              dayInWeekly: _selectedDaysNumber,
                              frequency: int.parse(_frequencyController.text),
                              resetOnMonth: _resetOnMonthCheckbox,
                            );

                            provider.createRoutine(newRoutine);
                            widget.fetchDataTrigger();

                          }
                        }
                      },
                      child: Text('Create'),
                      
                  );

                }else{

                  return ElevatedButton(
                    child: Text(isDeleted? "Revert" : "Delete"),
                    onPressed: () {
                      if(isDeleted){
                        provider.revertRoutine(widget.routine!.id!);
                      }else{
                        provider.deleteRoutine(widget.routine!.id!);
                      }
                      setState(() {
                        isDeleted = !isDeleted;
                      });
                    }
                  );

                }
              }
            )

          ],
        ),
      ),
    );
  }
}
