using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    class HotKey
    {
        public Guid Id { get; }
        public string _description { get; }
        private string[] Combination;

        public HotKey(Guid id, string description, string[] combination)
        {
            Id = id;
            _description = description;
            Combination = combination;
        }

        public IEnumerable<string> _combination => Combination.AsEnumerable();

    }
}
