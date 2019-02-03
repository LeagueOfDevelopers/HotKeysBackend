using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    public class Key
    {
        int Code;
        string Name;

        public Key(int code, string name)
        {
            Code = code;
            Name = name;
        }

        public override bool Equals(object obj)
        {
            var key = obj as Key;
            return key != null &&
                   Code == key.Code &&
                   Name == key.Name;
        }

        public override int GetHashCode()
        {
            var hashCode = -168117446;
            hashCode = hashCode * -1521134295 + Code.GetHashCode();
            hashCode = hashCode * -1521134295 + EqualityComparer<string>.Default.GetHashCode(Name);
            return hashCode;
        }
    }
}
