﻿using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    class Program_
    {
        public Guid id { get; }
        public string name { get; }
        private List<HotKey> HotKeys;
        public IEnumerable<HotKey> hotKeys => HotKeys.AsEnumerable();
        public void AddNewHotKey(HotKey hotKey)
        {
            HotKeys.Add(hotKey);
        }
        public void ChangeHotKey(Guid hotKeyId, HotKey newHotKey)
        {
            var id = HotKeys.FindIndex(hotk => hotk.Id == hotKeyId);
            HotKeys[id] = newHotKey;
        }
    }
}
